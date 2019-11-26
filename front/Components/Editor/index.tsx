import React, { FunctionComponent, useState, useRef, } from 'react'
import styles from './index.less'
const rust = import('../../../pkg/index.js')

export const Editor: FunctionComponent = () => {
  const [markdownSource, setMarkdownSource] = useState('')
  const markdownSourceRef = useRef<HTMLTextAreaElement>(null)
  const markdownResultRef = useRef<HTMLDivElement>(null)
  const handleOnChange = (value: string) => {
    rust.then(module => {
      const result = module.parse(value)
      markdownResultRef.current.innerHTML = result
    })
  }
  return (
    <div className={styles['editor-wrapper']}>
      <div
        className={styles['markdown-side']}
      >
        <textarea
          className={styles.textarea}
          ref={markdownSourceRef}
          value={markdownSource}
          onChange={evt => {
            setMarkdownSource(evt.target.value)
            handleOnChange(evt.target.value)
          }} />
      </div>
      <div
        className={styles['preview-side']}
        ref={markdownResultRef}
      ></div>
    </div>
  )
}