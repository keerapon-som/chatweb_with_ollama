"use client";

import React, { useState } from "react";

export default function Generate() {
  const [messages, setMessages] = useState<{ text: string; sender: string }[]>(
    []
  );
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (input.trim()) {
      setMessages([...messages, { text: input, sender: "user" }]);
      setInput("");
      // Simulate a response from the bot
      setTimeout(() => {
        setMessages((prevMessages) => [
          ...prevMessages,
          { text: "This is a response from the bot", sender: "bot" },
        ]);
      }, 1000);
    }
  };

  return (
    <div className="w-screen bg-neutral-800 text-gray-400 h-screen ">
      <div className="flex items-center justify-center h-4/5 overflow-auto">
        <div>
          {messages.map((message, index) => (
            <div key={index} className="">
              {message.text}
            </div>
          ))}
        </div>
      </div>
      <div className="input-box h-1/5 bg-slate-900">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Type a message..."
        />
        <button onClick={handleSend}>Send</button>
      </div>
    </div>
  );
}
