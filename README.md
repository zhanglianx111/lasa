## Jenkins RESTful API
---
![jenkinsLogo](https://wiki.jenkins-ci.org/download/attachments/2916393/master-jenkins-normal3.svg)  
<h2 id="0">目录</h2>
* [1 JOB](#1)
 * [1.1 get all jobs information](#1.1)
 * [1.2 get a job information](#1.2)
 * [1.3 is job running?](#1.3)
 * [1.4 create a job](#1.4)
 * [1.5 delete a job](#1.5)
 * [1.6 disable a job](#1.6)
 * [1.7 enable a job](#1.7)
 * [1.8 rename a job](#1.8)
 * [1.9 get all builds of a job](#1.9)
 * [1.10 build a job](#1.10)
 * [1.11 get build log of a job](#1.11)
 * [1.12 stop build of a job](#1.12)
 * [1.13 copy job](#1.13)
* [2 NODE](#2)
 * [2.1 add a node](#2.1)
 * [2.2 delete a node](#2.2)
 * [2.3 get all nodes informat](#2.3)
 * [2.4 get a node information](#2.4)
 * [2.5 get a node idle info](#2.5)
 * [2.6 set a node online](#2.6)
 * [2.7 set a node offline](#2.7)
 * [2.8 is node online?](#2.8)
 * [2.9 get node launching log](#2.9)
 * [2.10 disconnect node](#2.10)
 
* [3 VIEW](#3)
 * [3.1 get all views](#3.1)
 * [3.2 get a view](#3.2)
 * [3.3 add a view](#3.3)
 * [3.4 add a job into view](#3.4)
 * [3.5 delete a job from view](#3.5)
* [4 LOG](#4)
 * [4.1 get log level](#4.1)
 * [4.2 set log level](#4.2)
* [5 PLUGINS](#5)
 * [5.1 get all plugins](#5.1)

### <h2 id="1">1.JOB</h2>
<h3 id="1.1">**1.1 get all jobs information** </h3>   
**Get** */api/job/alljobs*  
**Response Code :** (Status 200)
**response:**(json)  
```json
[
    {
        "buildable": bool,
        "builds": [
            {
                "number": int,
                "url": string
            },
        ],
        "color": string,
        "concurrentBuild": bool,
        "description": string,
        "displayname": string,
        "inQueue": bool,
        "lastBuild": {
            "number": int,
            "url": string
        },
        "lastFailedBuild": {
            "number": int,
            "url": string
        },
        "lastStableBuild": {
            "number": int,
            "url": string
        },
        "lastSuccessfulBuild": {
            "number": int,
            "url": string
        },
        "name": string,
        "url": string
    }
]
```

<h3 id="1.2">**1.2 get a job information**</h3>  
**Get** */api/job/ajob/{jobname}*   
*jobname*: job的名字  
**Response Code :** (Status 200)
**response:**(json)  
```json
{
    "buildable": bool,
    "builds": [
        {
            "number": int,
            "url": string
        },
    ],
    "color": string,
    "concurrentBuild": bool,
    "description": string,
    "displayname": string,
    "inQueue": bool,
    "lastBuild": {
        "number": int,
        "url": string
    },
    "lastFailedBuild": {
        "number": int,
        "url": string
    },
    "lastStableBuild": {
        "number": int,
        "url": string
    },
    "lastSuccessfulBuild": {
        "number": int,
        "url": string
    },
    "name": string,
    "url": string
}
```
<h3 id="1.3">**1.3 is job running?**</h3>  
 **Get** */api/job/isrunning/{jobname}*  
 *jobname*: job的名字  
 **Response Code :** (Status 200)  
 **response:**(bool)
```json
bool
```

<h3 id="1.4">**1.4 create a job**</h3>  
 **Post** */api/job/create/{jobname}*  
 *jobname*: job的名字  
 *请求body字段及含义
```json
{
    "description": string,
    "build": {
        "dockerbuildandpublish":{
            "repositryname": string, //用户自定义的repo name
            "tag": string,             //images tag infomation
            "dockerregitstry": string, //default: http://dhub.yunpro.cn
            "skippush": bool //是否跳过push到docker registry
        },
        "executeshell": {
            command: string
        }
    }
}  
```
 **Response Code :** (Status 200)
 ```json
{
    "name": string,
    "description": string,
    "url": string,
    "build": {
        "dockerbuildandpublish": {
            "repositryname": string,
            "tag": string,
            "dockerregistry": string,
            "skippush": bool
         }
         "executeshell": {
             "command": string
         }
    }
}
```

<h3 id="1.5">**1.5 delete a job**</h3>  
**Post** */api/job/delete/{jobname}*  
*jobname*: job的名字  
**Response Code :** (Status 200)
```json
{
    "name": string
}
```

<h3 id="1.6">**1.6 disable a job**</h3>  
**Post** */api/job/disable/{jobname}*  
*jobname*: job的名字  
**Response Code :** (Status 200)
```json
bool

```

<h3 id="1.7">**1.7 enable a job**</h3>  
**Post** */api/job/enable/{jobname}*  
*jobname*: job的名字  
**Response Code :** (Status 200)
```json
bool
```

<h3 id="1.8">**1.8 rename a job**</h3>  
**Post** */api/job/rename/{oldname}/{newname}*  
*oldname*: job的老名字
*newname*: job的新名字  
**Response Code :** (Status 200)
 ```json
 string
 ```

<h3 id="1.9">**1.9 get all builds of a job**</h3>  
**Get** */api/job/{jobname}/allbuilds*  
*jobname*: job的名字  
**Response Code :** (Status 200)
```json
[
   {
       "number": int,
       "url": string
   },
]
 ```

<h3 id="1.10">**1.10 build a job**</h3>  
**Post** */api/job/build/{jobname}*  
*jobname*: job的名字  
**Response Code :** (Status 200)
```json
bool
```

<h3 id="1.11">**1.11 get build log of a job**</h3>  
**Post** */api/job/buildlog/{jobname}/{buildnumber}*     
*jobname*: job的名字  
*buildnumber*： build id  
**Response Code :** (Status 200)  
```json
strings
```

<h3 id="1.12">**1.12 stop build of a job**</h3>  
**Post** */api/job/stopbuild/{jobname}/{buildnumber}*  
*jobname*: job的名字  
*buildnumber*： build id  
**Response Code :** (Status 200)  
```json
bool
```
<h3 id="1.13">**1.13 copy job from anothor one**</h3>  
**Post** */api/job/copy/{from}/{newname}*  
*from*: job的名字  
*newname*: 新job的名字  
**Response Code :** (Status 200)  
```json
bool
```
### <h2 id="2">2 NODE</h2>
<h3 id="2.1">**2.1 add a node**</h3>  
**Post** */api/node/create/{nodename}/{numexecutors}/{description}/{remotefs}*  
**Request Body**  
```json
    "method"                string `json:"method"`,
    "host"                  string `json:"host"`,
    "port"                  string `json:"port"`,
    "credentialsId"         string `json:"credentialsId"`,
    "jvmOptions"            string `json:"jvmOptions"`,
    "javaPath"              string `json:"javaPath"`,
    "prefixStartSlaveCmd"   string `json:"prefixStartSlaveCmd"`,
    "suffixStartSlaveCmd"   string `json:"suffixStartSlaveCmd"`,
    "maxNumRetries"         string `json:"maxNumRetries"`,
    "retryWaitTime"         string `json:"retryWaitTime"`,
    "lanuchTimeoutSeconds"  string `json:"lanuchTimeoutSeconds"`,
```
**Response Code :** (Status 200)  
```json
bool
```

<h3 id="2.2">**2.2 delete a node**</h3>  
**Post** */api/node/delete/{nodename}*  
**Response Code :** (Status 200) 
```json
bool
```

<h3 id="2.3">**2.3 get all nodes informat**</h3>
**Get** */api/node/allnodes*  
**Response Code :** (Status 200)  
```json
[
    {
        "name": string
    }
]
```

<h3 id="2.4">**2.4 get a node information**</h3>  
**Get** */api/node/anode/{nodename}*  
**Response Code :** (Status 200)  
```json
{
    "displayName": string,
    "idle": bool,
    "numExecutors": int,
    "offline": bool
}
```

<h3 id="2.5">**2.5 get a node idle info**</h3>  
**Get** */api/node/isidle/{nodename}*  
**Response Code :** (Status 200)  
```json
bool
```
 
<h3 id="2.6">**2.6 set a node online******</h3>  
**Post** */api/node/online/{nodename}*  
**Response Code :** (Status 200)  
```json
bool
```

<h3 id="2.7">**2.7 set a node offline**</h3>   
**Post** */api/node/offline/{nodename}*  
**Response Code :** (Status 200)  
```json
bool
```
<h3 id="2.8">**2.8 is node online?**</h3>  
**Get** */api/node/isonline/{nodename}*  
**Response Code :** (Status 200)  
```json
bool
```

<h3 id="2.9">**2.9 get node launching log**</h3>  
**Get** */api/node/log/{nodename}*  
**Response Code :** (Status 200)  
```json
string
```

<h3 id="2.10">**2.10 disconnect node**</h3>  
**Post** */api/node/disconnect/{nodename}*  
**Response Code :** (Status 200)  
```json
bool
```

### <h2 id="3">3 VIEW</h2>
<h3 id="3.1">**3.1 get all views**</h3>  
**Get** */api/view/allviews*  
**Response Code :** (Status 200) 
```json
[
    {
        "description": string,
        "jobs": [
            {
                "name": string,
                "url": string,
                "color": string
            },
            {
                "name": string,
                "url": string,
                "color": string
            }
        ],
        "name": string,
        "url": string
    }
]
```

<h3 id="3.2">**3.2 get a view**</h3>  
**Get** */api/view/{viewname}*  
**Response Code :** (Status 200) 
```json
{
    "description": string,
    "jobs": [
        {
            "name": string,
            "url": string,
            "color": string
        },
        {
            "name": string,
            "url": string,
            "color": string
        }
    ],
    "name": string,
    "url": string
}
```

<h3 id="3.3">**3.3 add a view**</h3>  
**Post** */api/view/create/{viewname}/{type}*  
**Response Code :** (Status 200) 
```json
bool
```

<h3 id="3.4">**3.4 add a job into view**</h3>  
**Post** */api/view/addjob/{viewname}/{jobname}*  
**Response Code :** (Status 200) 
```json
bool
```

<h3 id="3.5">**3.5 delete a job from view**</h3>  
**Post** */api/view/deletejob/{viewname}/{jobname}*  
**Response Code :** (Status 200) 
```json
bool
```

### <h2 id="4">4 LOG</h2>
<h3 id="4.1">**4.1 get log level**</h3>  
**Get** */api/log/level*  
**Response Code :** (Status 200)  
```json
string
```

<h3 id="4.2">**4.2 set log level**</h3>  
**Post** */api/log/set/{level}*  
**Response Code :** (Status 200)  
```json
string
```
###<h2 id="5">5 PLUGINS</h2>

<h3 id="5.1">**5.1 get all plugins**</h3>  
**Get** */api/plugins/{depth}*  
*depth*:  
- 0: 获取所有plugins    
- 1:   

**Response Code :** (Status 200)   
```json
[
    {
        "active": bool,
        "enabled": bool,
        "name": string,
        "url": string,
        "version": string
    },
]
```

[回到顶部](#0)
