{ pkgs ? import <nixpkgs> {} }:

let
  tex = (pkgs.texlive.combine {
    inherit (pkgs.texlive) scheme-medium
      wrapfig amsmath ulem hyperref capt-of csquotes fvextra
      upquote
      collection-langcyrillic cm-super babel-russian cyrillic;
  });
in
pkgs.mkShell {

  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.delve
    pkgs.pandoc
    pkgs.poppler_utils
    tex

    # keep this line if you use bash
    pkgs.bashInteractive
  ];
  shellHook = ''
    set -a
    source ${./.env}
    set +a

    export CGO_ENABLED=0
  '';
}
