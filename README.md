These things might help:

```
export CGO_ENABLED=1
```

Build with:
```
xk6 build --with github.com/Kristina-Pianykh/pubsubext@latest
```

Run with:
```
K6_ENABLE_COMMUNITY_EXTENSIONS=true ./k6 run script.js
