'use client';

import { useState } from 'react';
import { api } from '@/lib/api';

export default function TestAPI() {
  const [result, setResult] = useState<string>('');
  const [loading, setLoading] = useState(false);

  const testConnection = async () => {
    setLoading(true);
    setResult('Testing connection...');
    
    try {
      const response = await api.get('/users');
      setResult(`✅ Success! Found ${response.data.length} users`);
    } catch (error: unknown) {
      const errorMessage = error instanceof Error ? error.message : 'Unknown error';
      setResult(`❌ Error: ${errorMessage}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4">
      <div className="max-w-md mx-auto bg-white rounded-lg shadow p-6">
        <h1 className="text-xl font-bold mb-4">API Connection Test</h1>
        
        <button
          onClick={testConnection}
          disabled={loading}
          className="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 disabled:opacity-50"
        >
          {loading ? 'Testing...' : 'Test API Connection'}
        </button>
        
        {result && (
          <div className="mt-4 p-3 bg-gray-100 rounded">
            <pre className="text-sm">{result}</pre>
          </div>
        )}
        
        <div className="mt-4 text-sm text-gray-700">
          <p><strong>การทดสอบการเชื่อมต่อ API:</strong></p>
          <p className="font-mono text-blue-600">http://localhost:8080/api/v1/users</p>
          <p className="mt-2 text-amber-600">⚠️ ต้องเปิด Go backend ให้ทำงานก่อน!</p>
        </div>
      </div>
    </div>
  );
}