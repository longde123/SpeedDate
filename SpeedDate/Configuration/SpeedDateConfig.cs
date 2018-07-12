﻿using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Reflection;
using System.Runtime.Serialization;
using NullGuard;

namespace SpeedDate.Configuration
{
    public sealed class SpeedDateConfig
    {
        private readonly List<IConfig> _pluginConfigs = new List<IConfig>();

        public NetworkConfig Network { get; internal set; }

        public PluginsConfig Plugins { get; internal set; }

        public void Add(IConfig config)
        {
            _pluginConfigs.Add(config);
        }

        
        public bool TryGetConfig(string typeName, [AllowNull] out IConfig result)
        {
            result = _pluginConfigs.FirstOrDefault(config => config.GetType().FullName.Equals(typeName));
            return result != null;
        }
    }

    public class PluginsConfig
    {
        public static readonly PluginsConfig LoadAllPlugins = new PluginsConfig();
        public static readonly PluginsConfig DefaultPeerPlugins = new PluginsConfig("SpeedDate.ClientPlugins;SpeedDate.ClientPlugins.Peer*");
        public static readonly PluginsConfig DefaultSpawnerPlugins = new PluginsConfig("SpeedDate.ClientPlugins;SpeedDate.ClientPlugins.Spawner*");
        public static readonly PluginsConfig DefaultGameServerPlugins = new PluginsConfig("SpeedDate.ClientPlugins;SpeedDate.ClientPlugins.GameServer*");

        public static readonly PluginsConfig DefaultServerPlugins = new PluginsConfig("SpeedDate.ServerPlugins*");

        public string Namespaces { get; set; }

        public PluginsConfig(string namespaces = "*")
        {
            Namespaces = namespaces;
        }
    }

    public class NetworkConfig
    {
        [AllowNull]
        public string Address { get; set; } = String.Empty;
        public int Port { get; set; }

        public string Key { get; set; }

        public override string ToString()
        {
            return $"{Address}:{Port}";
        }

        public NetworkConfig(string address = default(string), int port = default(int), string key = "TundraNet")
        {
            Address = address;
            Port = port;
            Key = key;
        }

        public NetworkConfig(IPAddress address, int port = default(int), string key = "TundraNet")
        {
            Address = address.ToString();
            Port = port;
            Key = key;
        }
    }
}
