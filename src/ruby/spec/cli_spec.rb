require "cli"

RSpec.describe CLI do
  it "parses the input file and writes to the output stream" do
    args = ["../../example-data/survey.csv", "participation"]

    output = StringIO.new

    expected_output = <<~EOF
      Participation

      Participants: 6
      Submitted: 5 (83.3%)
    EOF

    CLI.new.call(arguments: args, out: output)

    expect(output.read).to eq(expected_output)
  end

  xit "does not show participation data when given that argument" do
    args = ["../../example-data/survey.csv", "average_results"]

    output = StringIO.new

    expected_output = ""

    CLI.new.call(arguments: args, out: output)

    expect(output.read).to eq(expected_output)
  end

  xit "does not show participation data when given that argument " do
    args = ["../../example-data/survey.csv", "average_frequency"]

    output = StringIO.new

    expected_output = ""

    CLI.new.call(arguments: args, out: output)

    expect(output.read).to eq(expected_output)
  end
end
