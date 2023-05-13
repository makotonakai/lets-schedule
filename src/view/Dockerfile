FROM node:18

WORKDIR /app
COPY ./ /app

RUN npm install -g npm && \
    npm install -g vue-cli

EXPOSE 3000

CMD ["npm", "run", "dev" ]
