class CfPurge < Formula
    desc "Purge CloudFormation stacks that match a glob pattern"
    homepage "https://github.com/dzhuneyt/cf-purge"
    url "https://github.com/Dzhuneyt/cf-purge/archive/refs/tags/0.0.6.tar.gz"
    sha256 "b7109b6e813d710e52568e7201d1442744830bda6bb711d9932bc8e076a82bd6"
    license "MIT"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", *std_go_args
    end
  
    test do
      system "#{bin}/cf-purge", "--version"
    end
  end