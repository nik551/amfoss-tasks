require 'nokogiri'
require 'open-uri'
itis =Nokogiri::HTML(open("https://www.google.com/search?start=1&end=10&q=linux"))
itis.xpath("//a/div").each do |link|
puts  link.text
puts " "
end