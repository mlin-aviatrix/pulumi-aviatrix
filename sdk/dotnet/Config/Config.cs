// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Aviatrix
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("aviatrix");

        private static readonly __Value<string?> _controllerIp = new __Value<string?>(() => __config.Get("controllerIp"));
        public static string? ControllerIp
        {
            get => _controllerIp.Get();
            set => _controllerIp.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<string?> _pathToCaCertificate = new __Value<string?>(() => __config.Get("pathToCaCertificate"));
        public static string? PathToCaCertificate
        {
            get => _pathToCaCertificate.Get();
            set => _pathToCaCertificate.Set(value);
        }

        private static readonly __Value<bool?> _skipVersionValidation = new __Value<bool?>(() => __config.GetBoolean("skipVersionValidation") ?? true);
        public static bool? SkipVersionValidation
        {
            get => _skipVersionValidation.Get();
            set => _skipVersionValidation.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username"));
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

        private static readonly __Value<bool?> _verifySslCertificate = new __Value<bool?>(() => __config.GetBoolean("verifySslCertificate"));
        public static bool? VerifySslCertificate
        {
            get => _verifySslCertificate.Get();
            set => _verifySslCertificate.Set(value);
        }

    }
}
