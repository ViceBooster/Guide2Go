# This is a Fork of the broken EPGO, from the often missing in action chuchodavids.

He changed the name of the project and readme, but not the fuckin' $PATH for his new EPGO commands, so here's a cleaner readme.

### Docker Install (via Docker Compose File)

I need to set up the docker image. However, you can build your own by running:

**docker-compose**
```
version: "3.4"
services:
    guide2go:
      container_name: guide2go
      image: chuchodavids/guide2go:stable
      ports:
        - 8080:8080
      environment:
        - TZ: America/Chicago
      volumes:
        - /YOU_APP_PATH/guide2go:/app
        - /YOUR_IMAGE_PATH:/app/images
      restart: always
```

## Using the APP

```guide2go -h```

```bash
-config string
    = Get data from Schedules Direct with configuration file. [filename.yaml]
-configure string
    = Create or modify the configuration file. [filename.yaml]
-h  : Show help
```

### Create a config file

**note: You can use the sample config file that is in the /config folder inside of the docker container**

```guide2go -configure MY_CONFIG_FILE.yaml```  
If the configuration file does not exist, a YAML configuration file is created. 

**Configuration file from version 1.0.6 or earlier is not compatible.**  

#### Terminal Output

```txt
Configuration [MY_CONFIG_FILE.yaml]
-----------------------------
 1. Schedules Direct Account
 2. Add Lineup
 3. Remove Lineup
 4. Manage Channels
 5. Create XMLTV File [MY_CONFIG_FILE.xml]
 0. Exit
```

##### Follow the instructions in the terminal

1. Schedules Direct Account:  
Manage Schedules Direct credentials.  

2. Add Lineup:  
Add Lineup into the Schedules Direct account.  

3. Remove Lineup:  
Remove Lineup from the Schedules Direct account.  

4. Manage Channels:  
Selection of the channels to be used.
All selected channels are merged into one XML file when the XMLTV file is created.
When using all channels from all lineups it is recommended to create a separate guide2go configuration file for each lineup.  
5. Create XMLTV File [MY_CONFIG_FILE.xml]:  
Creates the XMLTV file with the selected channels.
