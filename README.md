# Zone4-Fight-District-Server-Manager  
## Channel Manager (Local Backend & Web Panel)

> **Note:** Although it is written and contains some references to "MU", this is only a disguise.

A system for **controlling Channel Server execution locally**, now with a web panel for easy management via your browser.

---

## ⚠️ Important: Requires the Server Main

To use the Channel Manager, you **must also run the Main Server**:  
- [Zone4-Server-Main-GO](https://github.com/Tsujiu/Zone4-Server-Main-GO)  
- The Main Server is essential for the full Zone4 experience and handles game logic.

---

## Structure

- `config/channels.json`: List of channels and run commands
- `controllers/`: API handlers
- `process/`: Manages the actual processes (Start / Stop)
- `main.go`: Entry point for the API server
- `web/`: Web panel for controlling channels via browser (localhost)

---

## How to Use

1. **Place this folder next to your `mu-server/` directory.**
2. **Edit `config/channels.json`** if you want to add or change channels.
3. **Start the backend:**
    ```
    go run main.go
    ```
4. **Access the Web Panel:**  
    Open your browser and go to:  
    [http://localhost:8080/](http://localhost:8080/)  
    - Login with your password.
    - Start, stop, restart channels and view status and logs visually.

5. **API (Optional, for advanced users):**  
    Use the API via Postman, curl, or your own integrations:
    ```
    POST http://localhost:8080/start/channel1
    POST http://localhost:8080/stop/channel1
    GET  http://localhost:8080/status
    POST http://localhost:8080/start-all
    POST http://localhost:8080/stop-all
    ```

---

## Notes

- **Primarily for Windows** (`cmd /C ...` commands)
- **Unlimited** number of channels supported
- Web panel is for **local use** (localhost) and protected by password
- Project references "MU" only as a disguise; real focus is Zone4
