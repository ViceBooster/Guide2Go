# This is a Fork of the broken EPGO, from the often missing in action chuchodavids.

I implemented changes that were missing in the current epgo version.

**Recent changes:**

1. For players like tiviMate I added the "live" and "new" icon in the tittle of the program. This is usually auto-added in Emby or other IPTV players, but TiviMate does not do this by reading XML tags.
1. Searching for channel to add is no longer a painful task. It can be done easily from the terminal
1. Add images to shows from The Movie Database

### Docker Install (via Docker Compose File)

I need to set up the docker image. However, you can build your own by running:

**docker-compose**
```
version: "3.4"
services:
    epgo:
      container_name: guide2go
      image: chuchodavids/epgo:stable
      ports:
        - 8080:8080
      environment:
        - TZ: America/Chicago
      volumes:
        - /YOU_APP_PATH/epgo:/app
        - /YOUR_IMAGE_PATH:/app/images
      restart: always
```

## Using the APP

```epgo -h```

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
When using all channels from all lineups it is recommended to create a separate epgo configuration file for each lineup.  
5. Create XMLTV File [MY_CONFIG_FILE.xml]:  
Creates the XMLTV file with the selected channels.
