import { useState, type SubmitEvent } from 'react'
import './App.css'

type Status = {
  type: "success" | "error";
  message: string;
};

function App() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [message, setMessage] = useState("");
    const [loading, setLoading] = useState(false);
    const [status, setStatus] = useState<Status | null>(null);

    async function handleSubmit(event: SubmitEvent<HTMLFormElement>) {
        event.preventDefault();
        setLoading(true);
        setStatus(null);

        try {
            const response = await fetch("/api/v1/contact", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ name, email, message }),
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || "Failed to send message.");
            }

            setStatus({
                type: "success",
                message: data.message || "Message sent successfully!",
            });

            setName("");
            setEmail("");
            setMessage("");
        } catch (error) {
            setStatus({
                type: "error",
                message:
                    error instanceof Error
                        ? error.message
                        : "Something went wrong.",
            });
        } finally {
            setLoading(false);
        }
    }

    return (
        <main className="container">
            <header>
                <span className="terminal-label">PortfolioSSH</span>
            </header>

            <form onSubmit={handleSubmit}>
                <label htmlFor="name">Name</label>
                <input
                    id="name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Your name"
                    required
                />

                <label htmlFor="email">Email</label>
                <input
                    id="email"
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="you@example.com"
                    required
                />

                <label htmlFor="message">Message</label>
                <textarea
                    id="message"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    placeholder="Send me a message pls"
                    rows={6}
                    required
                />

                <button type="submit" disabled={loading}>
                    {loading ? "Sending..." : "Send message →"}
                </button>

                {status && (
                    <p
                        className={`status ${status.type}`}
                        role="status"
                        aria-live="polite"
                    >
                        {status.message}
                    </p>
                )}
            </form>

            <footer>
                <span>Larp larp larp sahur</span>
            </footer>
        </main>
    );
}

export default App
