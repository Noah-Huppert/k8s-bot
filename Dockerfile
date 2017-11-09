FROM node:alpine

# Dependencies
COPY package.json .
COPY package-lock.json .
COPY node_modules node_modules

RUN npm install

# Rest of code
COPY external-scripts.json hubot-scripts.json ./
COPY scripts scripts
COPY bin bin

# Run
CMD npm start
