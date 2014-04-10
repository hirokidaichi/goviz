#!/usr/bin/env perl
use strict;
use warnings;

my @repos = (
    [q{github.com/hirokidaichi/goviz},q{-ignore-test}],
    [q{github.com/mattn/gof},q{-ignore-test}],
    [q{github.com/mattn/anko},q{-ignore-libs -ignore-test}],
    [q{github.com/hashicorp/serf},q{-ignore-libs -ignore-test}],
    [q{github.com/lestrrat/go-xslate},q{-ignore-libs -ignore-test }],
    [q{github.com/tsenart/vegeta},q{-ignore-test}],
    [q{github.com/mitchellh/packer},q{-ignore-libs ignore-test}],
    [q{github.com/dotcloud/docker/docker},q{-ignore-test -ignore-libs -level 1}],
    [q{github.com/dotcloud/docker/docker},q{-ignore-test -ignore-libs -level 2}],
    [q{github.com/dotcloud/docker/docker},q{-ignore-test -ignore-libs -level 3}],
    [q{github.com/dotcloud/docker/runtime/execdriver/execdrivers/},q{-ignore-libs -ignore-test}],
);

sub name_of_image {
    my $r = shift;
    my $o = shift;
    $r =~s/[\.\/]/_/g;
    $o =~s/\s+/_/g;
    return "./images/$r$o.png";
}

for my $t (@repos) {
    my ($r,$option ) = @$t;
    my $image = name_of_image($r,$option);
    my $command = "go run ./goviz.go -i $r  $option | dot -Tpng -o $image";
    print "$command\n";
    `$command`;
}


