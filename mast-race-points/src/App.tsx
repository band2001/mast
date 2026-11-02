import { createHashRouter, RouterProvider } from 'react-router-dom'
import { Home } from './components/Home'
import { RacerManagement } from './components/RacerManagement'
import { Layout } from './components/Layout'
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { Season } from './components/Season';
import { Race } from './components/Race';
import { RacePoints } from './components/RacePoints';

const theme = createTheme({
  components: {
    MuiButtonBase: {
      defaultProps: {
        disableRipple: true,
      },
    },
  },
});

const router = createHashRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      { path: '/', element: <Home /> },
      { path: '/racer-management', element: <RacerManagement /> },
      { path: '/season-management', element: <Season /> },
      { path: '/race-management', element: <Race />},
      { path: '/race-points', element: <RacePoints />}
    ]
  }
])

const App = () => <ThemeProvider theme={theme}>
    <RouterProvider router={router} />
</ThemeProvider>

export default App