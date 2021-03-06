[![Build Status](https://cloud.drone.io/api/badges/elahe-dastan/haproxy/status.svg)](https://cloud.drone.io/elahe-dastan/haproxy)
# haproxy
I have had a repository for learning and practicing 'nats' then I started thinking about what the difference is<br/>
between different technologies for load balancing.The biggest difference is that technologies like 'nats' and<br/>
'rabbitmq' are used to dispatch jobs inside a project to different workers but technologies like 'nginx' and 'haproxy'<br/>
are used to dispatch incoming http requests to different servers.
In this repository I want to learn how to use 'haproxy' to balance http traffic.

## How to use haproxy
Using haproxy is super simple regardless of the language of backend here I have an easy code in golang and then added <br/>
haproxy.<br/>
My code creates an endpoint and with every get request it just sleeps 5 seconds and then says hi.<br/>
![](image/echo.png)

![](image/handler.png)<br/><br/>
then we need haproxy itself, I use docker-compose to run it<br/><br/>
![](image/docker-compose.png)

Haproxy needs a config file called haproxy.cfg that I show later, the point I want to mention in here is that since<br/>
I have used docker-compose to run haproxy I vae to mount the haproxy.cfg file as shown in the screenshot

Here is my haproxy.cfg file<br/><br/>
![](image/haproxy.png)

Merge Request by [Parham Alvani](https://github.com/1995parham)