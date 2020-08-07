FROM nginx:1.13.6

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]