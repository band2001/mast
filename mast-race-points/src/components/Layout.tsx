import { Outlet } from "react-router-dom"
import { MainNavigation } from "./MainNavigation"
import { Paper } from "@mui/material"
import { COPY } from "../constants/copy"

export const Layout = () => {
    return (
        <div style={{ margin: '0px' }}>
            <Paper 
                sx={{ 
                    backgroundColor: 'var(--mast-blue)', 
                    paddingY: '0.1px', 
                    paddingLeft: '20px', 
                    borderTopLeftRadius:'0px', 
                    borderTopRightRadius: '0px' 
                }} 
                elevation={3}
            >
                <h1 style={{ color: 'var(--mast-white)' }}>{COPY.HEADING.TITLE}</h1>
            </Paper>
            <Outlet />
            <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
                <MainNavigation />
            </Paper>
        </div>
    )
}