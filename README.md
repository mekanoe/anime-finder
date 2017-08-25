# anime-finder

this is a WIP side-project for learning recommender systems.

## abstract

this tool pulls data from kitsu, then generates a graph and a few neural network models from the data in order to find the best anime titles for you.

research from the netflix prize made it very apparent that *one* method isn't adequate, so we do a few different methods at once. some is to prevent cold-start issues, but turn themselves off after enough data is gathered.

## building

you need 
- **Go** (ideally 1.9+) 
- **node.js** (ideally 8.4+)
- **docker-compose** 
- go **dep**

```bash
# get dependencies
dep ensure

# build the binaries
./tool install

# start dev services
docker-compose up -d

# set up environment vars
eval $(./tool env)

# currently, running the services is volatile,
# so best bet is to just run based on cmd folder names.
# as the project matures, this documentation will mature with it.
```

## misc

most of the work was streamed on twitch, http://twitch.tv/jumpystick

also consider joining my community's discord, https://discord.gg/Z5C8JZu