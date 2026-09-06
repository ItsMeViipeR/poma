{
  description = "Environnement de dev Go";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-26.05";

  outputs =
    { nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          git
          go
          gopls
          air
        ];

        shellHook = ''
          export GOBIN=$PWD/.bin
          export PATH=$GOBIN:$PATH

          mkdir -p $GOBIN
        '';
      };
    };
}
