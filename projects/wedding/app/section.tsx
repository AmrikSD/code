import React, { PropsWithChildren } from "react"
import styles from "./section.module.css"

export default function Section({name, children}: PropsWithChildren<{name: string}>){
    return (

        <section
            id={name} 
            className={[styles.section, styles[name]].join(' ')} 
            key={name} data-page={name} >

            <div className={styles.container}>
                {children}
            </div>

        </section>
    )
}
