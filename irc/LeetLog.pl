#!/usr/bin/perl
use strict;
use warnings;
 
use Irssi;
use vars qw($VERSION %IRSSI);

use POSIX 'strftime';
use Time::HiRes 'gettimeofday';
 
$VERSION = "0.1";
%IRSSI = (
  authors       => 	"Jo Emil Holen - JoMs",
	name        => 	"LeetLog",
	description => 	"Logger for #scene.no at 13:37",
    license     => 	"GPLv2"
);

=po
    LOG FORMAT
    ----------
    valid/invalid decided if an entry is valid or invalid.
    * Too early
    * Too late
    * Multiple entries
    Timestamp with time down to milliseconds
    Nick
=cut

Irssi::settings_add_str($IRSSI{"name"}, "LeetLog_file", "");

# Error messages
my $filedir_not_set = "Filedirectory is not defined - /set LeetLog_file";

# Used when I leet
sub myleet {
    my ($server, $msg, $target) = @_;
    
    leet($server->{nick}, $msg);
}

# Used when others leet
sub otherleet {
    my ($server, $msg, $nick, $address, $target) = @_;
    
    leet($nick, $msg);
}

# Write to log
sub leet {
    my ($nick, $msg) = @_;
    
    my ($t, $tt)=gettimeofday;
    my $ms=sprintf("%03d",$tt/1000);
	my $time = strftime("%m-%d-%Y %H:%M:%S",localtime($t)) . ":$ms";
    
    # Check if logfile is defined. If yes - execute
    if (Irssi::settings_get_str("LeetLog_file"))
    {
        my $valid = "invalid";
        
        if ($msg =~ /(?i)^\s*$/)
        {
            $valid = "valid";
            $msg = length($msg);
        }
        
        my $log = $time ." ". $valid ." ". $nick ." ". $msg ."\n";
        
        #Open file and write info to it
        open (my $fh, ">>", Irssi::settings_get_str("LeetLog_file"));
        print $fh $log;
        close $fh;
    } else {
        Irssi::print($filedir_not_set);
    }
}

##Insert array and function for handling multiple entries here



# Signals needed, and their function calls
Irssi::signal_add("message public", "otherleet");
Irssi::signal_add("message own_public", "myleet");

# Checking settings
if (!Irssi::settings_get_str("LeetLog_file"))
{
    Irssi::print($filedir_not_set);
}