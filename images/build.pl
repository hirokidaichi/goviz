#!/usr/bin/env perl
use strict;
use warnings;

my @repos = (
    [q{github.com/hirokidaichi/goviz},q{-l}],
    [q{github.com/mattn/gof},q{}],
    [q{github.com/mattn/anko},q{}],
    [q{github.com/hashicorp/serf},q{}],
    [q{github.com/lestrrat/go-xslate},q{}],
    [q{github.com/tsenart/vegeta},q{-l}],
    [q{github.com/mitchellh/packer},q{-seek-in SELF -l}],
    [q{github.com/dotcloud/docker/docker},q{-seek-in github.com/dotcloud/docker -d 1}],
    [q{github.com/dotcloud/docker/docker},q{-seek-in github.com/dotcloud/docker -d 2}],
    [q{github.com/dotcloud/docker/docker},q{-seek-in github.com/dotcloud/docker -d 3}],
    [q{github.com/dotcloud/docker/runtime/execdriver/execdrivers/},q{-seek-in github.com/dotcloud/docker}],
);

sub name_of_image {
    my $r = shift;
    my $o = shift;
    $r =~s/github\.com\///;
    $r =~s/[\.\/]/-/g;
    $o =~s/github\.com\///;
    $o =~s/[\s+\/]/-/g;

    return "./images/$r$o.png";
}

`rm ./images/*.png`;

for my $t (@repos) {
    my ($r,$option ) = @$t;
    my $image = name_of_image($r,$option);
    my $command = "go run ./goviz.go -i $r  $option | dot -Tpng -o $image";
    print "$command\n";
    `$command`;
}


