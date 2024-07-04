# Step 1: Build the Next.js app
FROM node:22.4-alpine3.19 AS builder

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app
COPY . .

# Build the app
RUN npm run build

# Step 2: Serve the app with Nginx
FROM nginx:stable-alpine

# Copy the built files from the previous step
COPY --from=builder /app/out /usr/share/nginx/html

# Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]

