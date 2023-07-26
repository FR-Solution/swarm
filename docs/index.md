Базовая информация
========================

Что такое Рой
-----------------

"Рой" - Host Based Firewall, распространяемым по лицензии Apache 2.0, который был создан для упрощения процесса настройки и поддержки правил сетевого трафика с использованием технологии nftables и увеличения надежности передаваемых данных в условиях нулевого доверия. Этот инструмент позволяет легко настроить и управлять правилами брандмауэра, что делает его идеальным решением для использования в крупных корпоративных сетях, а также в малых и средних предприятиях.

???+ info 
    При разработке системы учитывались принципы простоты и надежности. Это позволяет эффективно реализовать микросегментацию сети с помощью брандмауэра, установленного на управляемом узле, и использовать декларативный подход при написании сетевых правил.

Какие проблемы решает
-----------------
Проект был создан с учетом потребностей крупных компаний в финансовых отраслях, которые столкнулись с рядом проблем в областе сетовой изоляции. Среди этих проблем можно выделить:

- Неоднородность конфигурации межсетевого оборудования от разных производителей.
- Сложность и высокая стоимость точечной изоляции (ip to ip, team to team).
- Отсутствие единого декларативного подхода к конфигурации сетевых правил.
- Высокая стоимость оборудования.

Преимущества
-----------------
Использование Роя, работающего на уровне операционной системы Linux, обладает несколькими преимуществами по сравнению с железными брандмауэра от популярных производителей сетевого оборудования:

- Удобство управления: Управление и настройка Роя осуществляется с помощью инструментов управления операционной системы, что обеспечивает простоту и удобство в работе, а также сокращает затраты на обслуживание. Это значительно упрощает задачу администрирования и управления системой.

- Низкая стоимость: Рой предоставляет возможность использовать уже имеющееся оборудование, что позволяет снизить затраты на приобретение дополнительного оборудования. Это особенно полезно для небольших организаций или отделов, у которых ограниченный бюджет на IT-инфраструктуру.

- Гибкость: Рой обеспечивает возможность настройки политик безопасности на уровне отдельных приложений и сервисов, что позволяет более гибко контролировать защиту.

- Масштабируемость: Рой обладает возможностью легкого масштабирования на большом количестве серверов, без необходимости приобретения дополнительного оборудования и проведения сложной интеграции.

- Улучшенная защита: Рой обладает высоким уровнем безопасности, так как он работает на уровне операционной системы каждого устройства в сети, что позволяет обеспечить защиту на более глубоком уровне.