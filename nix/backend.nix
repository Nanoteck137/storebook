{ self }: 
{ config, lib, pkgs, ... }:
with lib; let
  cfg = config.services.storebook;

  storebookConfig = pkgs.writeText "config.toml" ''
    listen_addr = "${cfg.host}:${toString cfg.port}"
    data_dir = "${cfg.dataDir}"
    password = "${cfg.password}"
    jwt_secret = "${cfg.jwtSecret}"
  '';
in
{
  options.services.storebook = {
    enable = mkEnableOption "Enable the storebook service";

    port = mkOption {
      type = types.port;
      default = 5285;
      description = "port to listen on";
    };

    host = mkOption {
      type = types.str;
      default = "";
      description = "hostname or address to listen on";
    };

    dataDir = mkOption {
      type = types.path;
      default = "/var/lib/storebook";
      description = "path to the data directory";
    };

    username = mkOption {
      type = types.str;
      description = "username of the first user";
    };

    password = mkOption {
      type = types.str;
      description = "password of the login";
    };

    jwtSecret = mkOption {
      type = types.str;
      description = "jwt secret";
    };

    package = mkOption {
      type = types.package;
      default = self.packages.${pkgs.system}.backend;
      description = "package to use for this service (defaults to the one in the flake)";
    };

    user = mkOption {
      type = types.str;
      default = "storebook";
      description = "user to use for this service";
    };

    group = mkOption {
      type = types.str;
      default = "storebook";
      description = "group to use for this service";
    };

    openFirewall = mkOption {
      type = types.bool;
      default = false;
      description = "open the ports in the firewall";
    };
  };

  config = mkIf cfg.enable {
    systemd.services.storebook = {
      description = "storebook";
      wantedBy = [ "multi-user.target" ];
      after = [ "network.target" ];

      serviceConfig = mkMerge [
        {
          User = cfg.user;
          Group = cfg.group;

          ExecStart = "${cfg.package}/bin/storebook serve -c '${storebookConfig}'";

          Restart = "on-failure";
          RestartSec = "5s";

          PrivateTmp = true;
          ProtectHome = true;
          ProtectHostname = true;
          ProtectKernelLogs = true;
          ProtectKernelModules = true;
          ProtectKernelTunables = true;
          ProtectProc = "invisible";
          ProtectSystem = "strict";
          RestrictAddressFamilies = [ "AF_INET" "AF_INET6" "AF_UNIX" ];
          RestrictNamespaces = true;
          RestrictRealtime = true;
          RestrictSUIDSGID = true;
        }

        (mkIf (cfg.dataDir != "/var/lib/storebook") {
          ReadWritePaths = [ cfg.dataDir ];
        })

        (mkIf (cfg.dataDir == "/var/lib/storebook") {
          StateDirectory = "storebook";
        })
      ];
    };

    networking.firewall = lib.mkIf cfg.openFirewall {
      allowedTCPPorts = [ cfg.port ];
    };

    users.users = mkIf (cfg.user == "storebook") {
      storebook = {
        group = cfg.group;
        isSystemUser = true;
        home = "${cfg.dataDir}";
      };
    };

    users.groups = mkIf (cfg.group == "storebook") {
      storebook = {};
    };
  };
}
