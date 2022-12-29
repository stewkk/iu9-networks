{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gopls

    # keep this line if you use bash
    pkgs.bashInteractive
  ];

  shellHook = ''
    set -a
    source ${./.env}
    set +a
  '';
}
