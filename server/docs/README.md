# Odyssey Server

The Server consists of 4 parts
- Game Logic
- Networking
- Meta
- Admin

Game Logic is dedicated to maintaining the game simulation.  
It handles resolving player actions, game events, and sharing results.  
Data persistence is handled by this layer.

Networking handles the dedicated, real time, bidirectional communication between
clients and the server.

Meta handles providing asynchronous data and actions with the server.  
e.g. listing a player's characters for selection before connecting through
the Networking layer.

Admin handles administrative tasks.  
e.g. an API that allows admins to upload Map updates.

## Player Join Flow
Client: sends request to Meta with Id.  
Meta: returns character list, any other info needed to start playing.  
Client: requests connection with Id and data returned from meta.  
Networking: upgrades to websocket persistent connection.  
Game: puts character in world.  
Client <-> Netowrk <-> Game, Player playing.  

This way, Player objects are always full players and not in  
"connected but not playing" states.
