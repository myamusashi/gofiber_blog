{pkgs ? import <nixpkgs> {}}: let
  inherit (pkgs) buildGoModule fetchFromGitHub mkShell;

  blogWebsite = buildGoModule rec {
    pname = "gofiber_blog";
    version = "0.1.0";

    src = fetchFromGitHub {
      owner = "myamusashi";
      repo = pname;
      rev = "main";
      hash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="; # Ganti dengan hash aktual
    };

    vendorHash = "sha256-0000000000000000000000000000000000000000000="; # Ganti dengan 'lib.fakeHash' pertama kali, lalu ganti dengan hash aktual

    meta = {
      description = "Blog website menggunakan Go Fiber";
      homepage = "https://github.com/myamusashi/gofiber_blog";
      license = pkgs.lib.licenses.gpl3;
    };
  };
in {
  packages = {
    inherit blogWebsite;
  };

  devShell = mkShell {
    buildInputs = with pkgs; [
      go
      gopls
      gotools
      air # Untuk live reload (opsional)
    ];

    inputsFrom = [blogWebsite];
    shellHook = ''
      echo "Dev shell untuk ${blogWebsite.pname}"
    '';
  };
}
