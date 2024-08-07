import React from 'react'
import { MessageDto } from '../../@types/broker.type';
import { ReverseBinanceAggTradeDto } from '../../@types/binance.type';

interface IContentColumn {
    messages: MessageDto[];
}

const ContentColumn = ({messages}: IContentColumn) => {
    const contentData: ReverseBinanceAggTradeDto[] = messages.map(message => {
        console.log("- ", message.value)
        try{
            return JSON.parse(message.value)
        } catch(err){
            return {}
        }
    })

  return (
     <div className="Column">
      <div>"Content"</div>
      {contentData.map((content, index) =>
          <div key={`column-${index}`} className='value'>
            <div>{content.symbol}</div>
            <div>{content.price_change}</div>
          </div>
      )}
    </div>
  )
}

export default ContentColumn