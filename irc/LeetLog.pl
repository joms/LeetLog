#!/usr/bin/perl
use strict;
use warnings;
 
use Irssi;
use vars qw($VERSION %IRSSI);

use POSIX 'strftime';
use Time::HiRes 'gettimeofday';
 
$VERSION = "0.1";
%IRSSI = (
    authors     => 	"Jo Emil Holen - JoMs",
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
    Message or number of spaces
=cut

# Define setting(s)
Irssi::settings_add_str($IRSSI{"name"}, "LeetLog_file", "");

# Error messages
my $error_filedir_not_set = "Filedirectory is not defined - /set LeetLog_file";

# Used when I leet
sub myleet {
    my ($server, $msg, $target) = @_;
    
    leet($server->{nick}, $msg, $target);
}

# Used when others leet
sub otherleet {
    my ($server, $msg, $nick, $address, $target) = @_;
    
    leet($nick, $msg, $target);
}

# Variable for saving the date of the last round, and users that has posted today
my $lastround = "";
my @users;

# Debug variable
my $debug = 0;

=po
    Valid/Invalid reasons:
        0 = Valid
        1 = Before 13:37
        2 = Text in 13:37
        3 = Already entered
        4 = After 13:37
        5 = Empty string before 13:37
        6 = Empty string after 13:37
=cut

# Write to log
sub leet {
    my ($nick, $msg, $chan) = @_;
    
    my ($t, $tt)=gettimeofday;
    my $ms=sprintf("%03d",$tt/1000);
    my $date = strftime("%m/%d/%Y",localtime($t));
    my $time = strftime("%H:%M:%S",localtime($t)) . ".$ms";

    #Check if channel is #scene.no
    if ($chan eq "#scene.no" || $debug eq 1)
    {
    	# Check if time is within 13:35 and 13:40
    	if (strftime("%H", localtime($t)) == 13 && strftime("%M", localtime($t)) >= 35 && strftime("%M", localtime($t)) <= 40 || $debug eq 1)
    	{
    	    # Check if it's a new day, reset the userlist if it is
        	if ($lastround ne $date)
        	{
        	    @users=();
        	    $lastround = $date;
        	}
            
            # Check if logfile is defined. If yes - execute
            if (Irssi::settings_get_str("LeetLog_file"))
            {
                # Define the valid variable
                my $valid = 1;

                # Replaces msg with hex-value
                my $x = $msg;
                $x =~ s/(.)/sprintf("%x",ord($1))/eg;

                =po
                    2 : ^B
                    3 : ^C
                    f : ^O
                =cut
                if ($x eq 2 || $x eq 3 || $x eq "f")
                {
                    $msg = " ";
                }
                
                # Check if it's before :37
                if (strftime("%M", localtime($t)) < 37)
                {
                    # Check if string is empty, for miss
                    if ($msg =~ /^\s+$/i)
                    {
                        # Statistic over number of spaces
                        $msg = length($msg);
                        # Adding milliseconds for stats
                        $msg = $msg ." ". substr $time, -6, 6;
                        # Invalid because of miss
                        $valid = 5;
                    } else {
                        # Invalid as it's before :37
                        $valid = 1;
                    }
                } else {
                    # Check if it's after 37
                    if (strftime("%M", localtime($t)) > 37)
                    {
                        # Check if string is empty, for miss
                        if ($msg =~ /^\s+$/i)
                        {
                            # Statistic over number of spaces
                            $msg = length($msg);
                            # Adding milliseconds for stats
                            $msg = $msg ." ". substr $time, -6, 6;
                            # Invalid because of miss
                            $valid = 6;
                        } else {
                            # Invalid as it's after :37
                            $valid = 4;  
                        }
                    } else { 
                        # Check if string is empty
                        if ($msg =~ /^\s+$/)
                        {
                            # Check if user has already made an entry
                            foreach my $u (@users)
                            {
                                if ($u eq $nick)
                                {
                                    # Invalid because user already has made an entry
                                    $valid = 3;
                                }
                            }
                            
                            # Statistic over number of spaces
                            $msg = length($msg);
                            # Adding milliseconds for stats
                            $msg = $msg ." ". substr $time, -6, 6;
                        } else {
                            # Invalid because of text in 13:37
                            $valid = 2;   
                        }
                    }
                }
                
                # Message is valid
                if ($valid == 0)
                {
                    # Add user to list of entries
                    push(@users, $nick);
                }
                
                #Construct logline
                my $log = $date ."-". $time ." ". $valid ." ". $nick ." ". $msg ."\n";
                
                #Open file and write info to it
                open (my $fh, ">>", Irssi::settings_get_str("LeetLog_file"));
                print $fh $log;
                close $fh;
            } else {
                Irssi::print($error_filedir_not_set);
            }
    	}
    }
}

# Signals needed, and their function calls
Irssi::signal_add("message public", "otherleet");
Irssi::signal_add("message own_public", "myleet");

# Checking settings
if (!Irssi::settings_get_str("LeetLog_file"))
{
    Irssi::print($error_filedir_not_set);
}
