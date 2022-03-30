import { FaMoon, FaSun } from "react-icons/fa";
import { useState } from 'react';
import { useEffect } from 'react';
import { useRef } from 'react';

const ToggleButton = () => {

    //setup darkmode
    const [darkMode, setDark] = useState(false);
    const darkModeRef = useRef(darkMode);
    darkModeRef.current = darkMode;

    const setDarkMode = () => {
        document.documentElement.classList.contains('dark')
            ? document.documentElement.classList.remove('dark')
            : document.documentElement.classList.add('dark');

        return setDark(o => !o)
    }


    return (
        <div>
            {
                !darkMode
                    ? <FaSun onClick={() => setDarkMode()}  className="duration-300 transform motion-safe hover:scale-125 text-3xl cursor-pointer text-yellow-400" />
                    : <FaMoon onClick={() => setDarkMode()} className="duration-300 transform motion-safe hover:scale-125 text-3xl cursor-pointer text-indigo-400" />
            }
        </div>
    )

}
export default ToggleButton;