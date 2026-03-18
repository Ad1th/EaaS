class Eaas < Formula
  desc "Eventually As A Service"
  homepage "https://github.com/Ad1th/EaaS"
  url "https://github.com/Ad1th/EaaS.git", branch: "master"
  version "0.1.0"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args
  end

  test do
    output = shell_output("printf 'hello\neaas exit\n' | #{bin}/eaas")
    assert_match "Eventually.", output
  end
end
