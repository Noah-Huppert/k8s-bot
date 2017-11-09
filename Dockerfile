FROM node:alpine

# Dependencies
COPY package.json .
COPY package-lock.json .
COPY node_modules .

RUN npm install

# Rest of code
COPY . .

# Run
CMD npm start
