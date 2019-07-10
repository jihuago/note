## sshd_config文件

sshd_config是OpenSSH SSH服务器守护进程配置文件，路径在/etc/ssh/sshd_config

```
AcceptEnv
	指定客户端发送的哪些环境变量将会被传递到会话环境中。需要注意的是，有些环境变量可能会被用于绕过禁止用户使用的环境变量。由于这个原因，该指令应该小心使用。
AllowGroups
	默认允许所有组登陆
AllowUsers
    这个指令后面跟着一串用空格分隔的用户名列表(其中可以使用*和?通配符)。默认允许所有用户登录。如果使用了这个指令，那么将仅允许这些用户登录，而拒绝其他所有用户。如果指定了USER@HOST模式的用户，那么USER和HOST将同时被检查。
AuthorizedKeysFile
    存放该用户可以用来登录的RSA/DSA公钥。%h 表示用户的主目录，%u表示该用户的用户名。经过扩展之后的值必须要么是绝对路径，要么是相对于用户主目录的相对路径。默认值是".ssh/authrized_keys"
Banner
	将这个指令指定的文件中的内容在用户进行认证前显示给远程用户。这个特性仅能用于SSH-2，none表示禁用这个特性。

Ciphers
	指定SSH-2允许使用的加密算法。多个算法之间使用逗号分隔。可以使用的算法如下：
	 "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr",
             "3des-cbc", "arcfour128", "arcfour256", "arcfour", "blowfish-cbc", "cast128-cbc"
             
HostbasedAuthentication
	推荐使用默认值no禁止这种不安全的认证方式。
PermitRootLogin
	是否允许root登陆。建议使用no
PermitEmptyPasswords
	是否允许密码为空的用户远程登陆。
PasswordAuthentication
	是否允许使用基于密码的认证。建议改为no
Port
	指定sshd守护进程监听的端口号，默认为22。可以使用多条指令监听多个端口。
```

