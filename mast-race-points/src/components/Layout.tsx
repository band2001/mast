import { Outlet } from "react-router-dom"
import { MainNavigation } from "./MainNavigation"

export const Layout = () => {
    return (
        <div>
            <Outlet />
            <MainNavigation />
        </div>
    )
}