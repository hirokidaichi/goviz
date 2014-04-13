#!/usr/bin/env perl
use strict;
use warnings;
use Text::Xslate;
use File::Slurp qw/write_file/;

package Sample {
    use Class::Accessor "antlers";
    has title  => ( is => "rw" );
    has repo   => ( is => "rw" );
    has image  => ( is => "rw" );
    has option => ( is => "rw" );

    sub new {
        my ( $class, $options ) = @_;
        return bless $options, $class;
    }

    sub exec {
        my ($self) = @_;
        my $command =
          sprintf( "go run ./goviz.go -i %s %s | dot -Tpng -o images/%s\n",
            $self->repo, $self->option, $self->image );
        print $command;
        `$command`;
    }
}

my @SAMPLES = map { Sample->new($_) } (
    {
        title  => "anko",
        image  => "anko.png",
        repo   => "github.com/mattn/anko",
        option => "",
    },
    {
        title  => "serf",
        image  => "serf.png",
        repo   => "github.com/hashicorp/serf",
        option => "",
    },
    {
        title  => "go-xslate",
        image  => "xslate.png",
        repo   => "github.com/lestrrat/go-xslate",
        option => "",
    },
    {
        title  => "vegeta",
        image  => "vegeta.png",
        repo   => "github.com/tsenart/vegeta",
        option => "-l",
    },
    {
        title  => "packer",
        image  => "packer.png",
        repo   => "github.com/mitchellh/packer",
        option => "--search SELF -l",
    },
    {
        title  => "docker plot depth 1",
        image  => "docker1.png",
        repo   => "github.com/dotcloud/docker/docker",
        option => "-s github.com/dotcloud/docker -d 1",
    },
    {
        title  => "docker plot depth 2",
        image  => "docker2.png",
        repo   => "github.com/dotcloud/docker/docker",
        option => "-s github.com/dotcloud/docker -d 2",
    },
    {
        title  => "docker plot depth 3",
        image  => "docker3.png",
        repo   => "github.com/dotcloud/docker/docker",
        option => "-s github.com/dotcloud/docker -d 3",
    },
    {
        title  => "docker's execdrivers",
        image  => "docker-execdrivers.png",
        repo   => "github.com/dotcloud/docker/runtime/execdriver/execdrivers/",
        option => "-s github.com/dotcloud/docker",
    },
);

$_->exec for @SAMPLES;

my $own = Sample->new(
    {
        image  => "own.png",
        repo   => "github.com/hirokidaichi/goviz",
        option => "-s SELF"
    }
);
$own->exec;

my $template = Text::Xslate->new;
my $usage    = `go run ./goviz.go -h 2>&1`;
my $metrics  = `go run ./goviz.go -i github.com/dotcloud/docker/docker -m`;

print "create README.md \n";
my $result = $template->render( 'build/readme.tx',
    { usage => $usage, samples => \@SAMPLES, metrics => $metrics, own => $own }
);

write_file( "./README.md", $result );
