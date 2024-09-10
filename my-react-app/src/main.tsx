import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import ChatSeller from './ChatSeller.tsx'
import './ChatSeller.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ChatSeller />
  </StrictMode>,
)
