import { createHashRouter, RouterProvider, Outlet, Link } from 'react-router-dom'
import { Home } from './components/Home'
import { RacerManagement } from './components/RacerManagement'
import { Layout } from './components/Layout'

const router = createHashRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { path: '/', element: <Home /> },
      { path: '/racer-management', element: <RacerManagement /> }
    ]
  }
])

const App = () => <RouterProvider router={router} />

export default App