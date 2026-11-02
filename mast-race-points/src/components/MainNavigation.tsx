import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material"
import { Assignment, Functions, Home as HomeIcon, Person, SunnySnowing } from "@mui/icons-material";
import { Link } from "react-router-dom";
import { useState } from "react";

export const MainNavigation = () => {
    const [value, setValue] = useState(0)
    const onChange = ((_: React.SyntheticEvent<Element, Event>, newValue: any) => {
        setValue(newValue)
    }) 

    return(
        <Box>
            <BottomNavigation
                value={value}
                onChange={onChange}
                showLabels
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
                <BottomNavigationAction
                    label="Season Management"
                    icon={<SunnySnowing />}
                    component={Link}
                    to="season-management"
                />
                <BottomNavigationAction
                    label="Race Management"
                    icon={<Assignment />}
                    component={Link}
                    to="race-management"
                />
                <BottomNavigationAction
                    label="Race Points"
                    icon={<Functions />}
                    component={Link}
                    to="race-points"
                />
            </BottomNavigation>
        </Box>
    )
}