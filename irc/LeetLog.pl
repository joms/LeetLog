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

=po
    Valid/Invalid reasons:
        0 = Valid
        1 = Before 13:37
        2 = Text in 13:37
        3 = Already entered
        4 = After 13:37
=cut

##Insert array and function for handling multiple entries here
my $lastround = "";
my @users;


# Write to log
sub leet {
    my ($nick, $msg) = @_;
    
    my ($t, $tt)=gettimeofday;
    my $ms=sprintf("%03d",$tt/1000);
    my $date = strftime("%m-%d-%Y",localtime($t));
	my $time = strftime("%H:%M:%S",localtime($t)) . ":$ms";
	
	if ($lastround ne $date)
	{
	    @users=();
	    $lastround = $date;
	}
    
    # Check if logfile is defined. If yes - execute
    if (Irssi::settings_get_str("LeetLog_file"))
    {
        my $valid = 0;
        
        if ($msg =~ /(?i)^\s*$/)
        {
            my $invalid = 0;
            foreach my $u (@users)
            {
                if ($u eq $nick)
                {
                    $valid = 3;
                    $invalid = 1;
                }
            }
            
            if ($invalid == 0)
            {
                $valid = 0;
                $msg = length($msg);
                
                push(@users, $nick);
            }
        }
        
        my $log = $date ." ". $time ." ". $valid ." ". $nick ." ". $msg ."\n";
        
        #Open file and write info to it
        open (my $fh, ">>", Irssi::settings_get_str("LeetLog_file"));
        print $fh $log;
        close $fh;
    } else {
        Irssi::print($filedir_not_set);
    }
}


# Signals needed, and their function calls
Irssi::signal_add("message public", "otherleet");
Irssi::signal_add("message own_public", "myleet");

# Checking settings
if (!Irssi::settings_get_str("LeetLog_file"))
{
    Irssi::print($filedir_not_set);
}