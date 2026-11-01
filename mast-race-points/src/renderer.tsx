import './index.css';

import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';

// Find the root div from your HTML
const container = document.getElementById('root');

if (container) {
  const root = createRoot(container);
  // Render your root React component
  root.render(
     <React.StrictMode>
      <App />
     </React.StrictMode>
  );
}
