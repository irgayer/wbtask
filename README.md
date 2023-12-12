# Инструкция по запуску
```bash
go run main.go
```
Не знаю нужно ли устанавливать пакеты. 

# Комментарии к исправлениям
### XSS
После создания REST API и когда все нормальные компании поделились на front и back команды, XSS стал не так страшен лично для меня, но теперь это проблема front-end разработчиков.
По рекомендациям: использовать функции/методы для экранирования. В python есть замечательная функция:
```python
shlex.quote() // Встретилась на соревновании. Невозможно сломать(?).
```
### IDOR
После прочтения статьи на хабре, как-то вспомнил свои прошлые проекты и не везде была защита от "самого умного"
человека на свете :D. Должна быть проверка на access у пользователя. 
```csharp
// Должно быть что-то типа того, код написан на C#.
comment = _commentsRepository.GetById(id);
if (comment is null)
    throw new EntityNotFoundException(typeof(Comment), id);

if (comment.UserId != _authContext.UserId)
    throw new AccessDeniedException(...);
    
// Продолжение.
```
### SQLI
Современные ORM библиотеки сами экранируют запросы, но в данном случае, я СПЕЦИАЛЬНО закодил SQLI, чтобы показать, что это плохо.
```go
// https://gorm.io/docs/security.html
userInput := "jinzhu;drop table users;"
// safe, will be escaped
db.Where("name = ?", userInput).First(&user)

// SQL injection
db.Where(fmt.Sprintf("name = %v", userInput)).First(&user)
```

### OS Command Injection
Плохая идея вообще вызывать какие-то команды операционной системы, но если очень надо, то нужно экранировать входные данные.
### Path Traversal
Похоже на OS Command Injection, но тут нужно экранировать пути. 
### Brute Force
1) В проекте возвращается четкая ошибка "User does not exists" и "Incorrect password", это неправильно, но нужно было показать вектор. 
2) Blocklist/Whitelist (работа DevOps'ов. Удачи им!), ограничение запросов. 
# Дополнительные комментарии 
Не успел сделать проект по-нормальному из-за семейных обстоятельств. Надо было просить скинуть задание в понедельник.
1) Хотелось Domain, Application, Infrastructure, UI layers, но на Go опыта не имею.
2) По кривому Auth: добавил бы контекст AuthContext interface:
```csharp
// В C# это выглядит так.
// Он инициализируется на этапе запроса, считывая токен из заголовка.
// А еще там есть аттрибут Authorize, который проверяет, что пользователь авторизован.
public interface IAuthContext
{
    int UserId { get; }
    bool IsAuthenticated { get; }
} 

// Пример использования.
public class UpdateCommentRequest : IRequest<Unit>
{
    public int Id { get; set; }
    public string Text { get; set; }
}

public class UpdateCommentRequestHandler : IRequestHandler<UpdateCommentRequest, Unit>
{
    private readonly IAuthContext _authContext;
    // Использование Repository и UnitOfWork очень холиварная тема, готов долго с вами спорить.
    private readonly ICommentsRepository _commentsRepository;

    public UpdateCommentRequestHandler(IAuthContext authContext, ICommentsRepository commentsRepository)
    {
        _authContext = authContext;
        _commentsRepository = commentsRepository;
    }

    public async Task<Unit> Handle(UpdateCommentRequest request, CancellationToken cancellationToken)
    {
        var comment = await _commentsRepository.GetById(request.Id);
        if (comment is null)
            throw new EntityNotFoundException(typeof(Comment), request.Id);

        if (comment.UserId != _authContext.UserId)
            throw new AccessDeniedException(...);

        comment.Text = request.Text;
        await _commentsRepository.Update(comment);

        return Unit.Value;
    }
}

// В контроллере.
[Authorize]
[HttpPatch("{id}")]
public async Task<IActionResult> UpdateComment(int id, [FromBody] UpdateCommentRequest request)
{
    request.Id = id;
    await _mediator.Send(request);
    return Ok();
}
```
3) В целом задание показалось очень очевидным и скучным, которое не отражает реальность. Уже не стал фиксить уязвимости, потому что это очевидно исправляется.
4) Не делал тесты, но в реальном проекте они были бы обязательно (Unit, Functional, Integration всё как вы любите).
5) Ииии это мой первый проект на Go, поэтому не судите строго.

# P.S
- [x] Имею опыт участия в разработке микросервисных проектов с нуля по DDD.
- [x] Больше интересна архитектура программного обеспечения, но на эту стажировку записался из-за интереса, т.к. много друзей учится на Cyber Security.
- [x] Имею опыт участия на CTF.
- [x] Имею опыт участия на хакатонах, в том числе Blockchain.

# P.P.S Умею, практикую.
- [X] Jira, Confluence
- [x] Git proficient user
- [x] Docker

Более подробно в резюме. Thx.