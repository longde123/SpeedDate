﻿using SpeedDate.Configuration;

namespace SpeedDate.ServerPlugins.Database
{
    public class DatabaseConfig : IConfig
    {
        public bool CheckConnectionOnStartup { get; set; } = false;
        public string Host { get; set; } = "localhost";
        public int Port { get; set; } = 26257;
        public string Username { get; set; } = "root";
        public string Password { get; set; } = string.Empty;
        public string Database { get; set; } = "speeddate";
    }
}

