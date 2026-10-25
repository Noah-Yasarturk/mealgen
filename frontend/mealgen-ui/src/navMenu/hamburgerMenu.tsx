import { PiHamburgerLight } from "react-icons/pi";
import "../navMenu/hamburgerMenu.scss";
import { useEffect, useState } from "react";

const dbPingUrl = `${process.env.REACT_APP_BACKEND_HOST}/ping-db`

const HamburgerMenu = ({ style }) => {
    const [message, setMessage] = useState(null);

    useEffect(() => {
        console.log(dbPingUrl);

        fetch(dbPingUrl)
            .then((response) => response.json())
            .then((response) => setMessage(response.message))
            .catch((error) => console.error(error));
    }, [])

    return (
        <div id="piHamburgerIconContainer" style={{ ...style, color: message === "db up" ? "#391566" : "red" }}>
            <PiHamburgerLight />
        </div>
    )
}

export default HamburgerMenu;
