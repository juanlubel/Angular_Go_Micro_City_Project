FROM node

WORKDIR /app

COPY ./package*.json ./

RUN npm install
RUN npm install -g @angular/cli

COPY . .

EXPOSE 4500

CMD sleep 5 && npm start