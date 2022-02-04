# pull the official base image
FROM node:alpine
# set working direction
COPY  /frontend /app/
WORKDIR /app/
EXPOSE 3000
# install application dependencies
COPY /frontend/package.json ./
COPY /frontend/package-lock.json ./
RUN npm i
# start app
CMD ["npm", "start"]