import React, { useState } from "react"
import styles from "../app/section.module.css"

const Spinner = () => (
        <>spinner</>
    )

export default function Rsvp() {
    const [isLoading, setIsLoading] = useState(true)

    return <>
        {isLoading && <Spinner />}
        <iframe src="https://docs.google.com/forms/d/e/1FAIpQLScrFkUx_3O4K5dqS4DqHgOPdUM71e8QiQJEhSusMzHBjDhQMQ/viewform?embedded=true"
            width="80%"
            height="95%"
            className={styles.rsvpform}
            onLoad={() => setIsLoading(false)}
        >Loading…</iframe>
    </>
}
