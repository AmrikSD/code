import React, { PropsWithChildren, useState } from "react"
import styles from "./accordion.module.css"


const Accordion = ({title, children} : PropsWithChildren<{title: string}>) => {
    const [isActive, setIsActive] = useState(false)

    return (
        <div className={styles.accordion}>
            <div className={styles.title} onClick={() => setIsActive(!isActive)}>
                <h2>{title}</h2>
                <div>{isActive ? "-" : "+" }</div>
            </div>
            {isActive &&
                <div className={styles.content}>
                    {children}
                </div>
            }
        </div>
    )
}

export default Accordion;
