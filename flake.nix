{
	description = "";

	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs";
		flake-utils.url = "github:numtide/flake-utils";
	};

	outputs = { self, nixpkgs, flake-utils, }:
		flake-utils.lib.eachDefaultSystem (system: 
			let 
				overlays = [];
				lib = nixpkgs.lib;
				pkgs = import nixpkgs { inherit system overlays; };
			in
				{
					packages = {
					};
					# defaultPackage = example;
					devShell = pkgs.mkShell {
						packages = with pkgs; [ go gopls gore graphviz ];
						shellHook = ''
						'';
					};
				}
		);
}
