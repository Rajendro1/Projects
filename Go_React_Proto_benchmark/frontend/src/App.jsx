import React, { useState } from "react";
import * as pb from "./protos/user"; // Protobuf compiled JS

const endpoints = {
    json: "http://localhost:8080/rest",
    resp: "http://localhost:8080/resp",
    proto: "http://localhost:8080/proto",
};

async function measureFetch(label, url, parser) {
    const start = performance.now();
    const res = await fetch(url);
    const buffer = await res.arrayBuffer();
    const size = buffer.byteLength;
    const result = parser(buffer);
    const end = performance.now();
    return { label, ms: end - start, size, count: result.length };
}

function parseJSON(buffer) {
    return JSON.parse(new TextDecoder().decode(buffer));
}

function parseRESP(buffer) {
    const str = new TextDecoder().decode(buffer);
    const lines = str.split("\r\n").filter(Boolean);
    const users = [];
    for (let i = 0; i < lines.length; i += 8) {
        users.push({
            id: parseInt(lines[i + 2]),
            name: lines[i + 4],
            email: lines[i + 6],
            age: parseInt(lines[i + 8]),
        });
    }
    return users;
}

function parseProtobuf(buffer) {
    const UserList = pb.UserList;
    const decoded = UserList.decode(new Uint8Array(buffer));
    return decoded.users;
}

export default function App() {
    const [results, setResults] = useState([]);

    const runBenchmarks = async () => {
        const bench = [
            await measureFetch("JSON", endpoints.json, parseJSON),
            await measureFetch("RESP", endpoints.resp, parseRESP),
            await measureFetch("Protobuf", endpoints.proto, parseProtobuf),
        ];
        setResults(bench);
    };

    return (
        <div className="p-4">
            <h1 className="text-xl font-bold mb-2">Web API Benchmark</h1>
            <button onClick={runBenchmarks} className="bg-blue-600 text-white p-2 rounded">
                Run Benchmark
            </button>
            <table className="mt-4 border-collapse border border-gray-400 w-full">
                <thead>
                    <tr>
                        <th className="border border-gray-400 px-2 py-1">Format</th>
                        <th className="border border-gray-400 px-2 py-1">Time (ms)</th>
                        <th className="border border-gray-400 px-2 py-1">Size (bytes)</th>
                        <th className="border border-gray-400 px-2 py-1">User Count</th>
                    </tr>
                </thead>
                <tbody>
                    {results.map((r) => (
                        <tr key={r.label}>
                            <td className="border border-gray-400 px-2 py-1">{r.label}</td>
                            <td className="border border-gray-400 px-2 py-1">{r.ms.toFixed(2)}</td>
                            <td className="border border-gray-400 px-2 py-1">{r.size}</td>
                            <td className="border border-gray-400 px-2 py-1">{r.count}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
