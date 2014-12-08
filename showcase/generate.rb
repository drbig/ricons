#!/usr/bin/env ruby
# coding: utf-8
# See LICENSE.txt for licensing information.

require 'open-uri'
require 'json'

N = 4
ROOT = 'http://127.0.0.1:3232'
table = 'Generator|Size|' + 1.upto(N).to_a.join('|') + "\n" +\
        '---------|----|' + ('-|' * (N-1)) + '-' + "\n"

info = JSON.parse(open(ROOT + '/info.json').read, :symbolize_names => true)
info[:generators].each_pair do |k,v|
  [16, 100].each do |s|
    STDERR.puts "Fetching #{N} examples for #{v} of #{s}x#{s} size"
    table += "#{k}|#{s}x#{s}|"
    tmp = []
    1.upto(N) do |i|
      fn = "#{k}-#{s}x#{s}-#{i}.png"
      open(fn, 'wb') {|f| f.write(open(ROOT + "/#{k}/png/#{s}/#{s}").read) }
      tmp.push("![#{k} #{i}](https://raw.github.com/drbig/ricons/master/showcase/#{fn})")
    end
    table += tmp.join('|') + "\n"
  end
end

STDERR.puts "Markdown table:"
puts table
