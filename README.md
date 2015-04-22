# Squid_Golang_Redirector
Go语言实现Squid重定向器

# 目的： #
由于众所周知的原因，Google的公共库无法访问，导致访问一些国外站点卡在加载js或css，导致这个页面要等到js加载超时才能显示。
通过Squid的重定向器，自动把到Google的请求，以http-302 重定向到国内，来解决。

例如

http://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js

http-302 重定向到

http://ajax.useso.com/ajax/libs/jquery/1.7.2/jquery.min.js




# Squid重定向器设置 #

	#redirect_program 重定向器程序或脚本路径
	redirect_program /usr/share/squid/goredirect

	#redirect_children控制重定向器池的size。默认值是5 个进程
	redirect_children 20
	redirect_rewrites_host_header off
	redirector_bypass on

	#默认所有请求都会发到重定向其中，添加ACL只有Googleapis进入重定向器
	acl googlehost dstdom_regex -i -n \.googleapis.com
	redirector_access allow googlehost



**重定向器接口**（来自：**《Squid 中文权威指南》** 11章）

重定向器在其标准输入里，每次一行的接受来自squid 的数据。每行包括下列四个元素，
以空格分开：

1. **请求URI**
1. **客户IP 地址和完全可验证域名**
1. **用户名，通过RFC 1413 ident 或代理验证**
1. **HTTP 请求方式**

例如：

**http://www.example.com/page1.html 192.168.2.3/user.host.name jabroni GET**

请求 URI 取自客户请求，包括任何查询条件。然而，分段标记（例如#字符和随后的文本）被移除了。

第二个元素包含客户 IP 地址，和可选的完整可验证域名（FQDN）。假如激活了log_fqdn
指令或使用了srcdomain ACL 元素，FQDN 才会设置。尽管那样，FQDN 也许仍未知，因为
客户网络管理员没有在其DNS里正确的设置反向指针区域。假如squid 不知道客户的FQDN，
它用一个短横线(-)代替。例如：

http://www.example.com/page1.html 192.168.2.3/- jabroni GET

假如 squid 了解请求背后的用户名，客户ident 域才会设置。假如使用了代理验证，ident
ACL 元素，或激活了ident_lookup_access，这点才会发生。然而请记住，ident_lookup_access
指令不会导致squid 延缓请求处理。换句话说，假如你激活了该指令，但没有使用访问控制，
squid 在写往重定向进程时，也许仍不知道用户名。假如squid 不知道用户名，它显示一个
短横线(-)。例如：

http://www.example.com/page1.html 192.168.2.3/- - GET
Squid 从重定向进程里读回一个元素：URI。假如squid 读取一个空行，原始URI 保留
不变。

重定向程序永不退出，除非在标准输入里发生 end-of-file。假如重定向进程确实过早退
出，squid 在cache.log 里写一条警告信息：

WARNING: redirector #2 (FD 18) exited

假如 50%的重定向进程过早退出，squid 会以致命错误消息退出。


# 参考文档： #

1. [golang终端输出进度更新的代码](http://golanghome.com/post/607 "http://golanghome.com/post/607")
2. Squid 中文权威指南
3. [Python重定向器](http://eleveni386.7axu.com/posts/2013/10/13/squid-url-tiao-zhuan/)

