User
- id*
- username*
- email_address
- firstname
- lastname
- hashed_password
- salt
idx1: username
idx2: id

UserLike
- user_id
- image_id
- create_time

Image
- id*
- user_id*
- url
- desc
- likes
- downloads
- create_time*
idx1: create_time
idx2: user_id, create_time
idx3: id

ImageTag
- tag_id*
- image_id*
- create_time*
idx1: tag_id, create_time
idx2: image_id

Tag
- tag_id
- tag_name

/user/get
- Req
    - email
    - password
    - JWT token in header
- Resp
    - user
    - cookies

/user/create
- Req
    - email
    - first_name
    - last_name
    - username
    -  password
- Resp
    - user
    - cookies

/image/get
- Req
    - page_size
    - cursor
- Resp
    - []image
    - next_cursor

/image/get_by_tag
- Req
    - tag_name
    - page_size
    - cursor
- Resp
    - []image
    - next_cursor

/image/get_by_user
- Req
    - username
    - page_size
    - cursor
- Resp
    - []image
    - next_cursor

/image/get_user_likes
- Req
    - jwt
- Resp
    - []image
    - next_cursor

/image/like
- Req
    - image_id
- Resp

/image/download
- Req
    - image_id
- Resp

/images/get
- paginate by create_time in image_tab
- Get * from image_tab where create_time < ? ORDER BY create_time DESC LIMIT ?
- Get * from user_tab where id in ?

/images/get_by_user
- paginate by user_id, create_time in image_tab
- GET * from image_tab where user_id = ? and create_time < ? ORDER BY create_time DESC LIMIT ?
- GET * from user_tab where id = ?

/images/get_by_tag
- paginate by create_time in image_tag_tab
- GET * from image_tag_tab where tag_id = ? and create_time < ? ORDER BY create_time DESC LIMIT ?
- GET * from image_tab where id in ?
- GET * from user_tab where id in ?