// Connects Go and JS by bridging "http" between the two languages.
// Defines an object in JS that allows Golang to read and write buffers, as if
// connecting to a server.

class HTTPConnection {
    GoBuffer = new ArrayBuffer(2 * 1024 * 1024);
    GoBufferSize = 0;
    JSBuffer = new ArrayBuffer(2 * 1024 * 1024);
    JSBufferSize = 0;

    // This is the object that Golang will use to read and write buffers.
    Go = {
        Read: (n) => {
            // Read at most n bytes from the JS buffer.
            n = Math.min(n, this.JSBufferSize);

            const view = new Uint8Array(this.JSBuffer.slice(0, n));

            this.JSBuffer = this.JSBuffer.slice(n);
            this.JSBufferSize -= n;

            return view;
        },

        Write: (data) => {
            // Copy data to the end of the Go buffer.
            const view = new Uint8Array(this.GoBuffer);

            view.set(data, this.GoBufferSize);
            this.GoBufferSize += data.length;

            return data.length;
        }
    };

    Read = () => {
        // Read from the Go buffer.
        const n = this.GoBufferSize;
        const view = new Uint8Array(this.GoBuffer.slice(0, n));

        this.GoBuffer = this.GoBuffer.slice(n);
        this.GoBufferSize -= n;

        return view;
    }

    Write = (data) => {
        // Copy data to the end of the JS buffer.
        const view = new Uint8Array(this.JSBuffer);

        if (typeof data === "string") {
            data = new TextEncoder().encode(data);
        }

        view.set(data, this.JSBufferSize);
        this.JSBufferSize += data.length;

        return data.length;
    }
};
