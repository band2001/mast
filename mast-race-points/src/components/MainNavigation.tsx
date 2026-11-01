import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material"
import { Home as HomeIcon, Person } from "@mui/icons-material";
import { Link } from "react-router-dom";
import { useState } from "react";

export const MainNavigation = () => {
    const [value, setValue] = useState(0)
    const onChange = ((_: React.SyntheticEvent<Element, Event>, newValue: any) => {
        setValue(newValue)
    }) 

    return(
        <Box sx={{ width: 500 }}>
            <BottomNavigation
                value={value}
                onChange={onChange}
            >
                <BottomNavigationAction 
                    label="Home" 
                    icon={<HomeIcon />}
                    component={Link}
                    to="/"
                />
                <BottomNavigationAction 
                    label="Racer Management" 
                    icon={<Person />} 
                    component={Link}
                    to="/racer-management"
                />
            </BottomNavigation>
        </Box>
    )
}