require 'gmail'

#
# usage:
# ruby gmail-send.rb email@example.com name token
#

email_to = ARGV[0]
name     = ARGV[1]
token    = ARGV[2]

gmail = Gmail.connect('<LOGIN>', '<PASSWORD>')

gmail.deliver do
  to email_to
  subject "Email subject"
  text_part do
    body <<-BODY
      Text body
    BODY
  end
  html_part do
    content_type 'text/html; charset=UTF-8'
    body <<-BODY
      <p>HTML body</p>
    BODY
  end
end

gmail.logout
