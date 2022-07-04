require 'csv'

class CLI
  def call(arguments: ARGV, out: $stdout, err: $stderr)
    # Your implementation starts here

    headers = []
    rows = []

    CSV.foreach(arguments[0]) do |row|
      if headers.count == 0
        row.each do |header|
          headers.push(header)
        end
        next
      end

      data_item = DataItem.new(row)

      rows.push(data_item)
    end

    participants = rows.count
    submitions = 0

    rows.each do |item|
      if !item.submitted_at.nil?
        submitions += 1
      end
    end

    submitions_pc = (100 * (submitions/participants.to_f)).round(1)

    out.write "Participation\n\n"
    out.write "Participants: #{participants}\n"
    out.write "Submitted: #{submitions} (#{submitions_pc}%)\n"
    out.rewind
    return out
  end
end

class DataItem
  attr_accessor :email, :id, :submitted_at

  def initialize(input)
    # values = input.split(",")
    values = input # .split(",")

    @email = values[0]
    @id = values[1]
    @submitted_at = values[2]
  end

end


#
# out.write "Hello"
# out.rewind
# return out
